package com.example;

import java.util.*;
import java.io.File;
import java.net.URL;
import java.net.URLClassLoader;
import java.lang.reflect.Method;
import org.apache.commons.configuration2.INIConfiguration;
import org.apache.commons.configuration2.builder.FileBasedConfigurationBuilder;
import org.apache.commons.configuration2.builder.fluent.Parameters;
import org.apache.commons.configuration2.ex.ConfigurationException;

public class Main {
    static Method tfwords;
    static Method tffrequencies;
    static List<Object[]> word_freqs;
    static Object wordsPlugin;
    static Object frequenciesPlugin;

    public static void loadPlugins() throws Exception {
        Parameters params = new Parameters();
        FileBasedConfigurationBuilder<INIConfiguration> builder =
            new FileBasedConfigurationBuilder<>(INIConfiguration.class)
            .configure(params.fileBased().setFileName("config.ini"));

        try {
            INIConfiguration config = builder.getConfiguration();
            
            // Fetch classes from config file
            String wordsModule = config.getString("Plugins.words");
            File wordsFile = new File(wordsModule);
            File wordsParentFile = wordsFile.getParentFile();

            String frequenciesModule = config.getString("Plugins.frequencies");
            File frequenciesFile = new File(frequenciesModule);
            File frequenciesParentFile = frequenciesFile.getParentFile();

            URL[] urls = { wordsParentFile.toURI().toURL(), frequenciesParentFile.toURI().toURL() };
            
            URLClassLoader wordsLoader = new URLClassLoader(new URL[] { urls[0] });
            Class<?> loadedWordClass = wordsLoader.loadClass("Words2"); // Load the Class into Memory
            wordsPlugin = loadedWordClass.getDeclaredConstructor().newInstance(); // Instantiate the Class as an Object
            tfwords = wordsPlugin.getClass().getDeclaredMethod("extractWords", String.class); // create method
            
            URLClassLoader freqLoader = new URLClassLoader(new URL[] { urls[1] });
            Class<?> loadedFreqClass = freqLoader.loadClass("Frequencies2"); // Load the Class into Memory
            frequenciesPlugin = loadedFreqClass.getDeclaredConstructor().newInstance(); // Instantiate the Class as an Object
            tffrequencies = frequenciesPlugin.getClass().getDeclaredMethod("top25", List.class);
        } catch (ConfigurationException e) {
            e.printStackTrace();
        }
    }
    public static void main(String[] args) throws Exception {
        loadPlugins();
        List<String> result = (List<String>) tfwords.invoke(wordsPlugin, args[0]);
        word_freqs = (List<Object[]>) tffrequencies.invoke(frequenciesPlugin, result);

        for (Object[] word_freq : word_freqs) {
            System.out.println(word_freq[0] + " - " + word_freq[1]);
        }
    }
}