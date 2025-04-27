using UnityEngine;

public class DestroyObject : MonoBehaviour
{
    public int objectHP = 0; // 初期化、ゲーム開始時にAPIから値が上書き
    public int currentRespawnCount = 0;

    private bool isDead = false;

    private void OnTriggerEnter(Collider other)
    {
        if (isDead) return;

        if (other.CompareTag("Shell"))
        {
            objectHP -= 1;
            Destroy(other.gameObject);

            if (objectHP <= 0)
            {
                isDead = true;

                GameManager.Instance.OnEnemyDefeated(); // ★ GameManagerに倒された通知

                Destroy(this.gameObject);
            }
        }
    }
}
