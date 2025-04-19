using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class EnemyShotShell : MonoBehaviour
{
    public GameObject enemyShellPrefab;
    public float shotSpeed;
    public AudioClip shotSound;
    public float shotInterval = 1.5f; // 弾を打つ間隔（秒）

    private float timer = 0f;

    void Update()
    {
        timer += Time.deltaTime;

        if (timer >= shotInterval)
        {
            FireShell();
            timer = 0f; // タイマーリセット
        }
    }

    void FireShell()
    {
        GameObject enemyShell = Instantiate(enemyShellPrefab, transform.position, Quaternion.identity);

        Rigidbody enemyShellRb = enemyShell.GetComponent<Rigidbody>();
        enemyShellRb.AddForce(transform.forward * shotSpeed);

        AudioSource.PlayClipAtPoint(shotSound, transform.position);
        Destroy(enemyShell, 3.0f);
    }
}
