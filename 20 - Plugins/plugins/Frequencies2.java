import java.util.*;

public class Frequencies2 {
    static List<Object[]> word_freq = new ArrayList<>();
    static List<String> words = new ArrayList<>();
    static List<String> firstLetters = new ArrayList<>();
    static List<Object[]> letterCount= new ArrayList<>();

    public static void frequencies() {
        for (String word : words) {
            if (word.isEmpty()) continue;
            boolean found = false;
            for (int i = 0; i < word_freq.size(); i++) {
                Object[] pairs = word_freq.get(i);
                if (pairs[0].equals(word)) {
                    word_freq.get(i)[1] = (int) pairs[1] + 1;
                    firstLetters.add(word.substring(0, 1));
                    found = true;
                    break;
                }
            }
            if (!found) {
                word_freq.add(new Object[]{word, 1});
                firstLetters.add(word.substring(0, 1));
            }
        }
        for (String letter : new HashSet<>(firstLetters)) {
            letterCount.add(new Object[]{letter, Collections.frequency(firstLetters, letter)});
        }
    }

    public static void sort() {
        for (int i = 0; i < letterCount.size(); i++) {
            for (int n = letterCount.size() - 1; n > i; n--) {
                if ((int) letterCount.get(n)[1] > (int) letterCount.get(n - 1)[1]) {
                    Object[] temp = letterCount.get(n);
                    letterCount.set(n, letterCount.get(n - 1));
                    letterCount.set(n - 1, temp);
                }
            }
        }
    }

    public static List<Object[]> top25(List<String> inputWords) {
        words = inputWords;
        frequencies();
        sort();
        return letterCount;
    }
}