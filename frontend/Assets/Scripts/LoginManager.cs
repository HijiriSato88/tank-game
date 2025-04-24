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
    public TMP_Text errorMessageText;
    private bool isErrorVisible = false;

    private string loginUrl = "http://localhost:8080/login";

    void Start()
    {
        loginButton.onClick.AddListener(OnLoginButtonClick);
    }

    void Update()
    {
        if (isErrorVisible && Input.GetMouseButtonDown(0))
        {
            errorMessageText.text = "";
            isErrorVisible = false;
        }
    }

    void OnLoginButtonClick()
    {
        string username = usernameInput.text;
        string password = passwordInput.text;

        if (string.IsNullOrEmpty(username) || string.IsNullOrEmpty(password))
        {
            errorMessageText.text = "Please enter your username and password.";
            isErrorVisible = true;
            return;
        }

        errorMessageText.text = "";
        isErrorVisible = false;
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

            // セレクトシーンへ遷移
            UnityEngine.SceneManagement.SceneManager.LoadScene("Select");
        }
        else
        {
            errorMessageText.text = "Incorrect user name or password.";
            isErrorVisible = true;
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
