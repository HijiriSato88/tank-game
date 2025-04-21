using UnityEngine;
using UnityEngine.UI;
using TMPro; 
using UnityEngine.Networking;
using System.Collections;

public class LoginManager : MonoBehaviour
{
    public TMP_InputField usernameInput;
    public TMP_InputField passwordInput;
    public Button loginButton;

    private string loginUrl = "http://localhost:8080/login";

    void Start()
    {
        loginButton.onClick.AddListener(OnLoginButtonClick);
    }

    void OnLoginButtonClick()
    {
        string username = usernameInput.text;
        string password = passwordInput.text;

        StartCoroutine(LoginRequest(username, password));
    }

    IEnumerator LoginRequest(string username, string password)
    {
        string jsonData = JsonUtility.ToJson(new LoginData { username = username, password = password });

        UnityWebRequest request = new UnityWebRequest(loginUrl, "POST");
        byte[] bodyRaw = System.Text.Encoding.UTF8.GetBytes(jsonData);
        request.uploadHandler = new UploadHandlerRaw(bodyRaw);
        request.downloadHandler = new DownloadHandlerBuffer();
        request.SetRequestHeader("Content-Type", "application/json");

        yield return request.SendWebRequest();

        if (request.result == UnityWebRequest.Result.Success)
        {
            string responseText = request.downloadHandler.text;
            Debug.Log("Login success: " + responseText);

            // JSONからtokenを抽出（簡易的な例）
            string token = JsonUtility.FromJson<TokenResponse>(responseText).token;
            PlayerPrefs.SetString("token", token);

            // ゲームシーンへ遷移
            UnityEngine.SceneManagement.SceneManager.LoadScene("Main");
        }
        else
        {
            Debug.LogError("Login failed: " + request.error);
        }
    }

    [System.Serializable]
    public class LoginData
    {
        public string username;
        public string password;
    }

    [System.Serializable]
    public class TokenResponse
    {
        public string token;
    }
}
