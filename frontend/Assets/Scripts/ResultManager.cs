using UnityEngine;
using TMPro;
using UnityEngine.UI;
using UnityEngine.SceneManagement;

public class ResultManager : MonoBehaviour
{
    public static ResultManager Instance;

    public GameObject resultCanvas;
    public TMP_Text resultText;
    public Button backButton;

    private void Awake()
    {
        // シングルトンにしておく（シーン上に1つ）
        if (Instance == null) Instance = this;
        else Destroy(gameObject);

        // 起動時は非表示
        resultCanvas.SetActive(false);

        backButton.onClick.AddListener(() =>
        {
            SceneManager.LoadScene("Select");
        });
    }

    public void ShowResult(int score)
    {
        resultText.text = $"SCORE: {score}";
        resultCanvas.SetActive(true);
    }
}
