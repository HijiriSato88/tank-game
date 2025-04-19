using System.Collections;
using UnityEngine;

public class EnemyRespawner : MonoBehaviour
{
    public GameObject enemyPrefab;
    public float respawnDelay = 3f;

    public void RespawnEnemy(Vector3 spawnPosition)
    {
        StartCoroutine(RespawnCoroutine(spawnPosition));
    }

    private IEnumerator RespawnCoroutine(Vector3 position)
    {
        yield return new WaitForSeconds(respawnDelay);

        Instantiate(enemyPrefab, position, Quaternion.identity);
    }
}
