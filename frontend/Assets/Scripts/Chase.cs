using UnityEngine;
using UnityEngine.AI;

public class Chase : MonoBehaviour
{
    public GameObject target;
    private NavMeshAgent agent;

    void Start()
    {
        agent = GetComponent<NavMeshAgent>();

        // targetが未設定ならPlayerタグを持つオブジェクト
        if (target == null)
        {
            GameObject player = GameObject.FindWithTag("Player");
            if (player != null)
            {
                target = player;
            }
        }
    }

    void Update()
    {
        if (target != null && agent.isOnNavMesh)
        {
            agent.destination = target.transform.position;
            agent.speed = 4f; 
        }
    }
}
