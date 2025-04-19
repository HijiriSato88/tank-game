using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class DestroyObject : MonoBehaviour
{
    public int objectHP;

    private void OnTriggerEnter(Collider other)
{
    if (other.CompareTag("Shell"))
    {
        objectHP -= 1;
        Destroy(other.gameObject);

        if (objectHP <= 0)
        {
            // リスポーンマネージャーを探して復活を依頼する
            GameObject respawner = GameObject.Find("EnemyRespawner");
            if (respawner != null)
            {
                respawner.GetComponent<EnemyRespawner>().RespawnEnemy(transform.position);
            }

            Destroy(this.gameObject);
        }
    }
}

}
