package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"code.local/test/models"
	"code.local/test/repositories"
)

var ErrSignatureDoesNotMatch = errors.New("signature does not match")

type SignatureService struct {
	signatureRepo *repositories.SignatureRepository
}

func NewSignatureService(signatureRepo *repositories.SignatureRepository) *SignatureService {
	return &SignatureService{
		signatureRepo: signatureRepo,
	}
}

func (s *SignatureService) Close() error {
	return s.Close()
}

func (s *SignatureService) SignAnswers(userID string, questions, answers []string) (string, error) {
	combined := append(questions, answers...)
	hash := sha256.Sum256([]byte(
		strings.Join(combined, ";"),
	))
	signature := hex.EncodeToString(hash[:])

	data := models.Data{
		UserID:    userID,
		Questions: questions,
		Answers:   answers,
		Signature: signature,
		Timestamp: time.Now(),
	}

	err := s.signatureRepo.SaveSignature(userID, data)
	if err != nil {
		return "", err
	}

	return signature, nil
}

func (s *SignatureService) VerifySignature(userID, signature string) (models.Data, bool, error) {
	storedSignature, err := s.signatureRepo.GetSignature(userID)
	if err != nil {
		return models.Data{}, false, err
	}

	if storedSignature.Signature == signature {
		return storedSignature, true, nil
	}

	return models.Data{}, false, ErrSignatureDoesNotMatch
}
