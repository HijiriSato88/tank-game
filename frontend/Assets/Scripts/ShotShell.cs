using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ShotShell : MonoBehaviour
{
    public GameObject shellPrefab;
    public float shotSpeed;
    public AudioClip shotSound;

    void Update()
    {
        // Spaceキー
        if (Input.GetKeyDown(KeyCode.Space))
        {
            // 砲弾のプレハブをインスタンス化
            GameObject shell = Instantiate(shellPrefab, transform.position, Quaternion.identity);

            // 砲弾に付いているRigidbodyコンポーネントにアクセスする。
            Rigidbody shellRb = shell.GetComponent<Rigidbody>();

            // forward（青軸＝Z軸）の方向に力を加える。
            shellRb.AddForce(transform.forward * shotSpeed);

            // 発射した砲弾を３秒後に破壊する。
            // （重要な考え方）不要になった砲弾はメモリー上から削除すること。
            Destroy(shell, 3.0f);

            // 砲弾の発射音を出す。
            //AudioSource.PlayClipAtPoint(shotSound, transform.position);
        }
    }
}