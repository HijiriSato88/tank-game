using UnityEngine;
using UnityEngine.SceneManagement;
using TMPro;
using UnityEngine.UI;
using System.Collections; 
using UnityEngine.Networking;


public class SelectManager : MonoBehaviour
{
    public TMP_Text welcomeText;
    public Button gameStartButton;
    public Button rankingButton;

    private string meUrl = "http://localhost:8080/auth/me";

    void Start()
    {
        // ボタンにイベントを登録
        gameStartButton.onClick.AddListener(OnGameStartClicked);
        rankingButton.onClick.AddListener(OnRankingClicked);

        StartCoroutine(LoadUserInfo());
    }

    void OnGameStartClicked()
    {
        SceneManager.LoadScene("Main");
    }

    void OnRankingClicked()
    {
        SceneManager.LoadScene("Ranking");
    }

    IEnumerator LoadUserInfo()
    {
        string token = PlayerPrefs.GetString("token", "");
        if (string.IsNullOrEmpty(token))
        {
            welcomeText.text = "Not logged in.";
            yield break;
        }

        UnityWebRequest request = UnityWebRequest.Get(meUrl);
        request.SetRequestHeader("Authorization", "Bearer " + token);
        yield return request.SendWebRequest();

        if (request.result == UnityWebRequest.Result.Success)
        {
            string json = request.downloadHandler.text;
            MeResponse response = JsonUtility.FromJson<MeResponse>(json);
            welcomeText.text = "Hello " + response.username;
        }
        else
        {
            welcomeText.text = "Failed to load user.";
        }
    }

    [System.Serializable]
    public class MeResponse
    {
        public string username;
    }
}
