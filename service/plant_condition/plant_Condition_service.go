package plantCondition

import (
	"miniproject/entities"
	plantCondition "miniproject/repo/plant_condition"
)

func NewPlantConditionService(pr plantCondition.PlantConditionRepoInterface) *PlantConditionService {
	return &PlantConditionService{
		plantConditionRepoInterface: pr,
	}
}

type PlantConditionService struct {
	plantConditionRepoInterface plantCondition.PlantConditionRepoInterface
}

func (plantConditionService PlantConditionService) FindCondition(userID int) ([]entities.PlantCondition, error) {
	// Mengambil data kondisi tanaman dan data terkait Plant dan User
	plantConditions, err := plantConditionService.plantConditionRepoInterface.FindCondition(userID)
	if err != nil {
		return nil, err
	}

	for i, condition := range plantConditions {
		// Ambil data plant terkait dari repository
		plant, err := plantConditionService.plantConditionRepoInterface.FindPlantByIDAndUser(condition.PlantID, userID)
		if err != nil {
			return nil, err
		}

		// Mengupdate Plant dengan data User
		condition.Plant = plant
		plantConditions[i] = condition
	}

	return plantConditions, nil
}

func (plantConditionService PlantConditionService) FindConditionByID(conditionID, userID int) (entities.PlantCondition, error) {
	condition, err := plantConditionService.plantConditionRepoInterface.FindConditionByID(conditionID, userID)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return condition, nil
}

func (plantConditionService PlantConditionService) CreateCondition(condition entities.PlantCondition) (entities.PlantCondition, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.CreateCondition(condition)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) UpdateCondition(plant entities.PlantCondition) (entities.PlantCondition, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.UpdateCondition(plant)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) DeleteCondition(plant entities.PlantCondition) error {
	err := plantConditionService.plantConditionRepoInterface.DeleteCondition(plant)
	if err != nil {
		return err
	}
	return nil
}

func (plantConditionService PlantConditionService) GetPlantByIDAndUser(plantID, userID int) (entities.Plant, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.FindPlantByIDAndUser(plantID, userID)
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) CheckPlantId(plantID, userID int) error {
	err := plantConditionService.plantConditionRepoInterface.CheckPlantId(plantID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (plantConditionService PlantConditionService) FindByID(plantID int) (entities.PlantCondition, error) {
	condition, err := plantConditionService.plantConditionRepoInterface.FindByID(plantID)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return condition, nil
}

func (plantConditionService PlantConditionService) SplitResponse(plantData []entities.PlantCondition) ([]map[string]interface{}, []map[string]interface{}) {
	var plantInfo []map[string]interface{}
	var dataInfo []map[string]interface{}

	for _, condition := range plantData {
		// Urutkan objek: plant berada di atas data
		plantInfo = append(plantInfo, map[string]interface{}{
			"id":       condition.ID,
			"plant_id": condition.PlantID,
		})

		dataInfo = append(dataInfo, map[string]interface{}{
			"date":              condition.Date,
			"moisture_level":    condition.MoistureLevel,
			"sunlight_exposure": condition.SunlightExposure,
			"temperature":       condition.Temperature,
			"notes":             condition.Notes,
		})
	}

	// Gabungkan plantInfo dan dataInfo ke dalam format yang sesuai
	var result []map[string]interface{}
	for i := range plantInfo {
		result = append(result, map[string]interface{}{
			"alant_condition": plantInfo[i], // plant di atas
			"data":            dataInfo[i],  // data di bawah
		})
	}

	return result, nil
}

func (service *PlantConditionService) FindPlantByID(plantID, userID int) (entities.Plant, error) {
	return service.plantConditionRepoInterface.FindPlantByID(plantID, userID)
}
