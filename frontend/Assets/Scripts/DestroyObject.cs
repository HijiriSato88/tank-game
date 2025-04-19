using UnityEngine;

public class DestroyObject : MonoBehaviour
{
    public int objectHP = 3;
    public int currentRespawnCount = 0;

    private bool isDead = false;

    private void OnTriggerEnter(Collider other)
    {
        // 弾が複数フレームに渡って当たると、OnTriggerEnter が複数回発火され2体復活するケースを防ぐために以下のフラグで解決
        if (isDead) return;

        if (other.CompareTag("Shell"))
        {
            objectHP -= 1;
            Destroy(other.gameObject);

            if (objectHP <= 0)
            {
                isDead = true;

                GameObject respawner = GameObject.Find("EnemyRespawner");
                if (respawner != null)
                {
                    respawner.GetComponent<EnemyRespawner>().RespawnEnemy(transform.position, currentRespawnCount + 1);
                }

                Destroy(this.gameObject);
            }
        }
    }
}
