using System.Collections;
using UnityEngine;
using UnityEngine.Networking;
using System;

[Serializable]
public class EnemyData
{
    public int id;
    public string name;
    public int hp;
    public float move_speed;
    public int score;
}

public class EnemyDataFetcher : MonoBehaviour
{
    public static EnemyDataFetcher Instance { get; private set; }

    public EnemyData enemyData;
    public bool isLoaded = false;

    private void Awake()
    {
        if (Instance == null)
        {
            Instance = this;
            DontDestroyOnLoad(gameObject);
        }
        else
        {
            Destroy(gameObject);
        }
    }

    public void FetchEnemyData(string enemyName)
    {
        StartCoroutine(FetchEnemyCoroutine(enemyName));
    }

    private IEnumerator FetchEnemyCoroutine(string enemyName)
    {
        string url = $"http://localhost:8080/enemies/name?name={enemyName}";
        UnityWebRequest request = UnityWebRequest.Get(url);

        yield return request.SendWebRequest();

        if (request.result == UnityWebRequest.Result.Success)
        {
            enemyData = JsonUtility.FromJson<EnemyData>(request.downloadHandler.text);
            isLoaded = true; // データロード完了時にtrue
            Debug.Log($"Enemy Data Loaded: {enemyData.name}, HP: {enemyData.hp}, Speed: {enemyData.move_speed}, Score: {enemyData.score}");
        }
        else
        {
            Debug.LogError("Failed to fetch enemy data: " + request.error);
        }
    }
}
