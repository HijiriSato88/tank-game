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
            Debug.Log("done");
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
