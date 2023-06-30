import java.util.HashMap;
import java.util.Map;
import java.util.Random;

public class DictionarySearch {
    public static Map<String, Integer> createDictionary(int dictSize) {
        Map<String, Integer> testDict = new HashMap<>();
        String prefix = "abcdefghijklmnopqrstuvwxyz_";
        for (int i = 0; i < dictSize; i++) {
            String key = prefix + String.format("%03d", i);
            testDict.put(key, i);
        }
        return testDict;
    }
    public static void main(String[] args) {
        int dictSize = 1000;
        int[] searchTimes = { 50, 500, 5000, 50000 };

        Map<String, Integer> testDict = createDictionary(dictSize);

        Random random = new Random();

        for (int x = 0; x < 5; x++) {
            System.out.print("---------------------------\n");
            for (int searchTime : searchTimes) {
                long startTime = System.currentTimeMillis();
                for (int i = 0; i < searchTime; i++) {
                    String randomKey = (String) testDict.keySet().toArray()[random.nextInt(testDict.size())];
                    Integer result = testDict.get(randomKey);
                }
                long endTime = System.currentTimeMillis();
                double searchTimeMillis = endTime - startTime;
                System.out.print("Search time = " + String.format("%.1f", searchTimeMillis) + " ms\n");
            }
        }
    }
}
