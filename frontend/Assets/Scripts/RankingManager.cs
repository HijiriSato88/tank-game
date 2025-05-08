using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Networking;
using TMPro;

public class RankingManager : MonoBehaviour
{
    public GameObject rankingEntryPrefab;
    public Transform rankingContainer;

    private string rankingUrl = "http://localhost:8080/ranking?limit=100";

    [System.Serializable]
    public class RankingEntry
    {
        public string username;
        public int high_score;
        public int rank;
    }

    [System.Serializable]
    public class RankingEntryList
    {
        public List<RankingEntry> entries;
    }

    void Start()
    {
        StartCoroutine(LoadRanking());
    }

    public void OnBackButtonClicked()
    {
        UnityEngine.SceneManagement.SceneManager.LoadScene("Select");
    }

    IEnumerator LoadRanking()
    {
        UnityWebRequest request = UnityWebRequest.Get(rankingUrl);
        yield return request.SendWebRequest();

        if (request.result == UnityWebRequest.Result.Success)
        {
            string json = "{\"entries\":" + request.downloadHandler.text + "}";
            RankingEntryList list = JsonUtility.FromJson<RankingEntryList>(json);

            foreach (RankingEntry entry in list.entries)
            {
                GameObject obj = Instantiate(rankingEntryPrefab, rankingContainer);
                TMP_Text text = obj.GetComponent<TMP_Text>();
                text.text = $"{entry.rank} . {entry.username} : {entry.high_score}";
            }
        }
        else
        {
            Debug.LogError("Failed to load ranking: " + request.error);
        }
    }
}
