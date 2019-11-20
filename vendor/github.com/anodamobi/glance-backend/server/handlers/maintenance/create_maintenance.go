package maintenance

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"github.com/anodamobi/glance-backend/db/models"
	"github.com/anodamobi/glance-backend/server/handlers/errs"
	"github.com/anodamobi/glance-backend/server/handlers/notifications"
	"github.com/minio/minio-go"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

type CreateMaintenanceRequest struct {
	Images      []string `json:"images"`
	Description string   `json:"description"`
	UserID      int64    `json:"user_id"`
}

func (h *handler) CreateMaintenance(w http.ResponseWriter, r *http.Request) {
	var req CreateMaintenanceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create maintenance request")
		errs.BadRequest(w, err)
		return
	}

	uid, err := uniqueID()
	if err != nil {
		h.log.WithError(err).Error("failed to generate new unique id for maintenance request")
		errs.InternalError(w)
		return
	}

	bucketName := "maintenance-" + uid
	err = h.s3.MakeBucket(bucketName, "")
	if err != nil {
		h.log.WithError(err).WithField("unique_id", uid).Error("failed to create new bucket for maintenance request")
		errs.InternalError(w)
		return
	}

	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::%s/*"],"Sid": ""}]}`

	err = h.s3.SetBucketPolicy(bucketName, policy)
	if err != nil {
		h.log.WithError(err).WithField("unique_id", uid).Error("failed to update new bucket policy")
		errs.InternalError(w)
		return
	}

	var fileNames []string
	for _, image := range req.Images {
		fileNames = append(fileNames, bucketName+time.Now().String())
		_, err := h.s3.PutObject(bucketName, bucketName+time.Now().String(), bytes.NewReader([]byte(image)), int64(len(image)), minio.PutObjectOptions{})
		if err != nil {
			h.log.WithError(err).WithField("unique_id", uid).Error("failed to put new object into bucket")
			errs.InternalError(w)
			return
		}
	}

	// todo: integrate minio
	maintenance := models.Maintenance{
		Images:      fileNames,
		Description: req.Description,
		UserID:      req.UserID,
		CreatedAt:   time.Now(),
		UniqueID:    uid,
	}

	err = h.mntncDB.Insert(maintenance)
	if err != nil {
		h.log.WithError(err).Error("failed to insert new maintenance request into db")
		errs.InternalError(w)
		return
	}

	user, err := h.usersDB.Get(req.UserID)
	if err != nil {
		h.log.WithError(err).Error("failed to get user from db")
		errs.InternalError(w)
		return
	}

	err = notifications.NewNotifier(h.log, h.oneSignal).Notify("Maintenance request created!", "You have been created maintenance request!", []string{user.DeviceID})
	if err != nil {
		h.log.WithError(err).Error("failed to create push notification for create maintenance request")
		errs.InternalError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func uniqueID() (string, error) {
	b := make([]byte, 40)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	data := binary.BigEndian.Uint64(b)
	parsedData := cast.ToString(data)
	return parsedData[2:18], nil
}
