using System.Collections;
using UnityEngine;

public class EnemyRespawner : MonoBehaviour
{
    public GameObject enemyPrefab;
    public float respawnDelay = 3f;
    public int maxRespawns = 4;

    public void RespawnEnemy(Vector3 position, int respawnCount)
    {
        if (respawnCount <= maxRespawns)
        {
            StartCoroutine(RespawnCoroutine(position, respawnCount));
        }else
        {
            TankHealth tank = FindObjectOfType<TankHealth>();
            if (tank != null)
            {
                int enemiesDefeated = maxRespawns;
                int score = enemiesDefeated * 100 + tank.tankHP * 10;

                tank.StartCoroutine(tank.SendScoreToServer(score));
            }
        }
    }

    private IEnumerator RespawnCoroutine(Vector3 pos, int count)
    {
        yield return new WaitForSeconds(respawnDelay);

        GameObject newEnemy = Instantiate(enemyPrefab, pos, Quaternion.identity);

        // 復活回数を新しい敵に渡す
        DestroyObject destroyScript = newEnemy.GetComponent<DestroyObject>();
        if (destroyScript != null)
        {
            destroyScript.currentRespawnCount = count;
        }
    }
}
