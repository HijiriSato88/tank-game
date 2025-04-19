using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.AI;

public class Chase : MonoBehaviour
{
    public GameObject target;
    private NavMeshAgent agent;

    void Start()
    {
        agent = GetComponent<NavMeshAgent>();
    }

    void Update()
    {
        // もしターゲットが存在しなければ何もしない
        if (target == null) return;

        // ターゲットの位置を目的地に設定
        agent.destination = target.transform.position;
    }
}
