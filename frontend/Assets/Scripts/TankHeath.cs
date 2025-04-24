using UnityEngine;
using TMPro;
using System.Collections;

public class TankHealth : MonoBehaviour
{
    public TMP_Text HPLabel;
    public int tankHP;
    public EnemyRespawner respawner;

    void Start()
    {
        HPLabel.text = "HP:" + tankHP;
    }

    private void OnTriggerEnter(Collider other)
    {
        if (other.gameObject.tag == "EnemyShell")
        {
            tankHP -= 1;
            HPLabel.text = "HP:" + tankHP;
            Destroy(other.gameObject);

            if (tankHP <= 0)
            {
                HPLabel.text = "HP:0";
                int enemiesDefeated = GetLatestRespawnCount();
                int score = ScoreManager.CalculateScore(tankHP, enemiesDefeated);

                StartCoroutine(ScoreManager.SendScoreToServer(score, () =>
                {
                    ResultManager.Instance.ShowResult(score);
                    Destroy(gameObject);
                }));
            }
        }
    }

    int GetLatestRespawnCount()
    {
        DestroyObject[] enemies = FindObjectsOfType<DestroyObject>();
        int latest = 0;
        foreach (var e in enemies)
        {
            if (e.currentRespawnCount > latest)
                latest = e.currentRespawnCount;
        }
        return latest;
    }
}
