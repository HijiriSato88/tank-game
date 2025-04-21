using UnityEngine;
using TMPro;
using UnityEngine.Networking;
using System.Collections;

public class TankHealth : MonoBehaviour
{
    public TMP_Text HPLabel;
    public int tankHP;

    public EnemyRespawner respawner;

    void Start()
    {
        HPLabel.text = "HP:" + tankHP;
    }

    private void OnTriggerEnter(Collider other)
    {
        if (other.gameObject.tag == "EnemyShell")
        {
            tankHP -= 1;
            HPLabel.text = "HP:" + tankHP;
            Destroy(other.gameObject);

            if (tankHP <= 0)
            {
                HPLabel.text = "HP:0";
                CalculateAndSendScore();
            }
        }
    }

    void CalculateAndSendScore()
    {
        int enemiesDefeated = GetLatestRespawnCount();
        int score = enemiesDefeated * 100 + tankHP * 10;

        Debug.Log("Score: " + score);

        StartCoroutine(SendScoreToServer(score));
    }

    int GetLatestRespawnCount()
    {
        // 最後の敵の currentRespawnCount を探す
        DestroyObject[] enemies = FindObjectsOfType<DestroyObject>();
        int latest = 0;
        foreach (var e in enemies)
        {
            if (e.currentRespawnCount > latest)
                latest = e.currentRespawnCount;
        }
        return latest;
    }

    IEnumerator SendScoreToServer(int score)
    {
        string token = PlayerPrefs.GetString("token", "");

        if (string.IsNullOrEmpty(token)) yield break;

        string url = "http://localhost:8080/auth/score";

        ScoreData data = new ScoreData { score = score };
        string json = JsonUtility.ToJson(data);

        UnityWebRequest req = new UnityWebRequest(url, "POST");
        byte[] bodyRaw = System.Text.Encoding.UTF8.GetBytes(json);
        req.uploadHandler = new UploadHandlerRaw(bodyRaw);
        req.downloadHandler = new DownloadHandlerBuffer();
        req.SetRequestHeader("Content-Type", "application/json");
        req.SetRequestHeader("Authorization", "Bearer " + token);

        yield return req.SendWebRequest();

        if (req.responseCode == 200)
        {
            Debug.Log("スコア送信成功");
        }
        else
        {
            Debug.LogError($"スコア送信失敗 ({req.responseCode}): " + req.error);
        }
        Destroy(gameObject);
    }

    [System.Serializable]
    public class ScoreData
    {
        public int score;
    }
}
