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
                Destroy(this.gameObject);
            }
        }
    }
}
