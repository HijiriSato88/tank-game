using System.Collections;
using UnityEngine;
using UnityEngine.Networking;

public static class ScoreManager
{
    private static int totalScore = 0;

    [System.Serializable]
    private class ScoreData
    {
        public int score;
        public string event_slug; 
    }

    public static void AddEnemyScore(int enemyScore)
    {
        totalScore += enemyScore;
    }

    public static int CalculateScore(int tankHP, int enemiesDefeated)
    {
        return enemiesDefeated * 100 + tankHP * 10;
    }

    public static int GetTotalScore(int playerRemainingHP)
    {
        return totalScore + (playerRemainingHP * 10);
    }

    public static IEnumerator SendScoreToServer(int score, System.Action onSuccess = null, System.Action<string> onFailure = null)
    {
        string token = PlayerPrefs.GetString("token", "");
        if (string.IsNullOrEmpty(token))
        {
            onFailure?.Invoke("トークンがありません");
            yield break;
        }

        string url = "http://localhost:8080/auth/score";

        // 日付が2026/5/8より前かチェック
        var now = System.DateTime.Now;
        var cutoff = new System.DateTime(2026, 5, 8);

        ScoreData data = new ScoreData { score = score };
        if (now < cutoff)
        {
            data.event_slug = "test_event_yearly";
        }

        string json = JsonUtility.ToJson(data);

        UnityWebRequest req = new UnityWebRequest(url, "POST");
        byte[] bodyRaw = System.Text.Encoding.UTF8.GetBytes(json);
        req.uploadHandler = new UploadHandlerRaw(bodyRaw);
        req.downloadHandler = new DownloadHandlerBuffer();
        req.SetRequestHeader("Content-Type", "application/json");
        req.SetRequestHeader("Authorization", "Bearer " + token);

        yield return req.SendWebRequest();

        if (req.responseCode != 200)
        {
            Debug.LogError($"スコア送信失敗 ({req.responseCode}): {req.error}");
            onFailure?.Invoke(req.error);
        }
        else
        {
            Debug.Log("スコア送信成功");
            onSuccess?.Invoke();
        }
    }
}
