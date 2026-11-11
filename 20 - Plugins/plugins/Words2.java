import java.util.*;
import java.io.File;

public class Words2{
    static List<String> data = new ArrayList<>();
    static List<String> words = new ArrayList<>();

    public static void readFile(String path) throws Exception {
        File file = new File(path);
        Scanner scanner = new Scanner(file);
        StringBuilder stringBuilder = new StringBuilder();
        while (scanner.hasNextLine()) {
            stringBuilder.append(scanner.nextLine()).append("\n");
        }
        scanner.close();
        data.add(stringBuilder.toString());
    }
    public static void filterCharsAndNormalize() {
        String lowerCase = data.get(0).toLowerCase();
        String cleanBook = lowerCase.replaceAll("[^a-zA-Z\\s]", "");
        data.set(0, cleanBook);
    }
    public static void scan() {
        String[] splitWords = data.get(0).split("\\s+");
        for (String w : splitWords) {
            words.add(w);
        }
    }
    public static void removeStopWords() throws Exception {
        File stopWordsFile = new File("../stop_words.txt");
        Scanner scanner = new Scanner(stopWordsFile);
        StringBuilder stringBuilder = new StringBuilder();
        while (scanner.hasNextLine()) {
            stringBuilder.append(scanner.nextLine());
        }
        scanner.close();
        String[] stopWordsArray = stringBuilder.toString().split(",");
        Set<String> stopWordsSet = new HashSet<>(Arrays.asList(stopWordsArray));
        for (int i = 0; i < words.size(); i++) {
            String word = (String) words.get(i);
            if (stopWordsSet.contains(word) || word.length() <= 1 || !word.contains("z")) {
                words.remove(i);
                i--;
            }
        }
    }
    public static List<String> extractWords(String path) throws Exception {
        readFile(path);
        filterCharsAndNormalize();
        scan();
        removeStopWords();
        return words;
    }
}