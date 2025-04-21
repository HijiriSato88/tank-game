using UnityEngine;
using TMPro;

public class TankHealth : MonoBehaviour
{
    public TMP_Text HPLabel;

    public int tankHP;

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

            if (tankHP <= 0){
                Destroy(gameObject);
            }
        }
    }
}
