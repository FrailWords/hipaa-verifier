package hipaa

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"hipaa-verifier/types"
)

func ProtectData(c *fiber.Ctx, serviceRule *InterServiceRule) (any, error) {
	if serviceRule.RecordType == FhirRecord {
		dataToBeProtected := make([]*types.EhrRecord, 0)
		if err := c.BodyParser(&dataToBeProtected); err != nil {
			return nil, err
		}
		protectedData := ProtectDataRecords(dataToBeProtected)
		return protectedData, nil
	}
	return nil, errors.New("unknown record type for data protection")
}

func VerifyData(c *fiber.Ctx, serviceRule *InterServiceRule) (bool, error) {
	if serviceRule.RecordType == FhirRecord {
		dataToBeVerified := make([]types.EhrRecord, 0)
		if err := c.BodyParser(&dataToBeVerified); err != nil {
			return false, err
		}
		verified := VerifyDataRecords(dataToBeVerified)
		return verified, nil
	}
	return false, errors.New("unknown record type for data protection")
}
