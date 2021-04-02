package utils

import (
	"../configs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetAllResume(accessToken string) (ids []string, err error) {
	req, err := http.NewRequest("GET", configs.MINE_RESUME_URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 200 {
		return nil, errors.New(res["error"].(string))
	}

	resumesJson := res["items"].([]interface{})
	result := make([]string, len(resumesJson))

	for i, resume := range resumesJson {
		tmp := resume.(map[string]interface{})
		result[i] = tmp["id"].(string)
	}
	return result, nil
}

func PublishResume(resumeId string, accessToken string) (statusCode int, err error) {
	url := fmt.Sprintf(configs.PUBLIS_RESUME_URL, resumeId)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
