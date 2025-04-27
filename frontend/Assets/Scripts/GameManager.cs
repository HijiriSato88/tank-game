using UnityEngine;
using System.Collections.Generic;
using System.Collections;

public class GameManager : MonoBehaviour
{
    public static GameManager Instance { get; private set; }

    private Queue<string> enemyQueue = new Queue<string>();
    public EnemyRespawner respawner;

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

    void Start()
    {
        enemyQueue.Enqueue("EnemyA");
        enemyQueue.Enqueue("EnemyB");
        enemyQueue.Enqueue("EnemyC");

        SpawnNextEnemy();
    }

    public void SpawnNextEnemy()
    {
        if (enemyQueue.Count == 0)
        {
            Debug.Log("Game Clear!");
            // ここでクリア演出やシーン遷移など
            return;
        }

        string nextEnemyName = enemyQueue.Dequeue();
        StartCoroutine(LoadAndSpawn(nextEnemyName));
    }

    private IEnumerator LoadAndSpawn(string enemyName)
    {
        EnemyDataFetcher.Instance.isLoaded = false;
        EnemyDataFetcher.Instance.FetchEnemyData(enemyName);

        while (!EnemyDataFetcher.Instance.isLoaded)
        {
            yield return null;
        }

        respawner.StartRespawn(Vector3.zero); // スポーン位置を設定
    }

    public void OnEnemyDefeated()
    {
        SpawnNextEnemy();
    }
}
