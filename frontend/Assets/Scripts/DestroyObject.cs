using UnityEngine;

public class DestroyObject : MonoBehaviour
{
    public int objectHP = 0;
    public int currentRespawnCount = 0;
    public int enemyScore = 0;

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

                ScoreManager.AddEnemyScore(enemyScore);

                GameManager.Instance.OnEnemyDefeated();

                Destroy(this.gameObject);
            }
        }
    }
}
