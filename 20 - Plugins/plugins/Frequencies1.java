import java.util.*;

public class Frequencies1 {
    static List<Object[]> word_freq = new ArrayList<>();
    static List<String> words = new ArrayList<>();

    public static void frequencies() {
        for (String word : words) {
            boolean found = false;
            for (int i = 0; i < word_freq.size(); i++) {
                Object[] pairs = word_freq.get(i);
                if (pairs[0].equals(word)) {
                    word_freq.get(i)[1] = (int) pairs[1] + 1;
                    found = true;
                    break;
                }
            }
            if (!found) {
                word_freq.add(new Object[]{word, 1});
            }
        }
    }
    public static void sort() {
        for (int i = 0; i < word_freq.size(); i++) {
            for (int n = word_freq.size() - 1; n > i; n--) {
                if ((int) word_freq.get(n)[1] > (int) word_freq.get(n - 1)[1]) {
                    Object[] temp = word_freq.get(n);
                    word_freq.set(n, word_freq.get(n - 1));
                    word_freq.set(n - 1, temp);
                }
            }
        }
    }
    public static List<Object[]> top25(List<String> inputWords) {
        words = inputWords;
        frequencies();
        sort();
        List<Object[]> word_count = new ArrayList<>();
        for (int i = 0; i < 25 && i < word_freq.size(); i++) {
            word_count.add(word_freq.get(i));
        }
        return word_count;
    }
}