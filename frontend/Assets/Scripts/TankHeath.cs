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
                GameManager.Instance.OnPlayerDead(tankHP);
                Destroy(this.gameObject);
            }
        }
    }
}
