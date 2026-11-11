<h1>wk6maven: Plugin-Based Term Frequency Analyzer</h1>
<p>A Maven-based Java application that dynamically loads text-parsing and frequency-counting plugins using an INI configuration file and the Java Reflection API (<code>URLClassLoader</code>). Depending on the configured plugins, it extracts normal words or words containing the letter 'z', removes stop words, and tallies either the top 25 most frequent words or the frequency of starting letters.</p>

<h2>Prerequisites</h2>
<ul>
    <li>Java Development Kit (JDK) 8 or higher</li>
    <li>Apache Maven 3.6 or higher</li>
    <li>Input text file (e.g., <code>book.txt</code>) and stop words file (<code>stop_words.txt</code>)</li>
</ul>

<h2>Directory Structure</h2>
<p>Ensure your project is structured as follows so Maven and the dynamic class loader can locate the source files, resources, and external plugins:</p>
<pre><code>swe262/
├── stop_words.txt
└── wk6maven/
    ├── pom.xml
    ├── plugins/
    │   ├── Words1.java
    │   ├── Words2.java
    │   ├── Frequencies1.java
    │   └── Frequencies2.java
    └── src/
        └── main/
            ├── java/
            │   └── com/
            │       └── example/
            │           └── Main.java
            └── resources/
                ├── book.txt
                └── config.ini
</code></pre>

<h2>Configuration</h2>
<p>Specify which plugins to dynamically load by updating <code>src/main/resources/config.ini</code> with the target plugin paths:</p>
<pre><code>[Plugins]
words = plugins/Words1.class
frequencies = plugins/Frequencies1.class
</code></pre>

<h2>Compilation and Execution</h2>
<p>Navigate to your project root directory (<code>wk6maven</code>) in your terminal and run the following commands:</p>
<ol>
    <li><strong>Compile the external plugin classes:</strong>
        <pre><code>javac plugins/*.java</code></pre>
    </li>
    <li><strong>Compile the Maven project:</strong>
        <pre><code>mvn clean compile</code></pre>
    </li>
    <li><strong>Execute the program:</strong>
        <pre><code>mvn exec:java -Dexec.mainClass="com.example.Main" -Dexec.args="./src/main/resources/book.txt"</code></pre>
    </li>
</ol>