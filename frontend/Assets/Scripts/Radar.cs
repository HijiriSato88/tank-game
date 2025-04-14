using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Radar : MonoBehaviour
{
    public Transform target;

    // 他のコライダーに触れている間中実行される
    private void OnTriggerStay(Collider other)
    {
        // もしも他のオブジェクトに「Player」というTag（タグ）が付いてれば
        if (other.CompareTag("Player"))
        {
            // LookAt()メソッドは指定した方向にオブジェクトの向きを回転
            transform.root.LookAt(target);
        }
    }
}