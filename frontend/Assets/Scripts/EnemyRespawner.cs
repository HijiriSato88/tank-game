using UnityEngine;
using System.Collections;

public class EnemyRespawner : MonoBehaviour
{
    public GameObject enemyPrefab;

    public void StartRespawn(Vector3 position)
    {
        StartCoroutine(SpawnEnemyWhenReady(position));
    }

    private IEnumerator SpawnEnemyWhenReady(Vector3 position)
    {
        while (!EnemyDataFetcher.Instance.isLoaded)
        {
            yield return null;
        }

        GameObject newEnemy = Instantiate(enemyPrefab, position, Quaternion.identity);

        DestroyObject destroyScript = newEnemy.GetComponent<DestroyObject>();
        if (destroyScript != null)
        {
            destroyScript.objectHP = EnemyDataFetcher.Instance.enemyData.hp;

            var agent = newEnemy.GetComponent<UnityEngine.AI.NavMeshAgent>();
            if (agent != null)
            {
                agent.speed = EnemyDataFetcher.Instance.enemyData.move_speed;
            }
        }
    }
}
