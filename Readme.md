SERVICES
DOCUMENTS
Dagger - example.md
PREVIEW AS 
EXPORT AS 
SAVE TO 
IMPORT FROM 
DOCUMENT NAME
Dagger - example.md
MARKDOWNPREVIEWToggle Mode
  
<h1 class="code-line" data-line-start=0 data-line-end=1 ><a id="Dagger__example_0"></a>Dagger - example</h1>
<p class="has-line-data" data-line-start="2" data-line-end="3"><a href="https://dagger.io/"><img src="https://dagger.io/logo.svg" alt="N|Solid"></a></p>
<p class="has-line-data" data-line-start="5" data-line-end="6">This demo aims to explain in a practical and simple way the use of dagger.</p>
<ul>
<li class="has-line-data" data-line-start="8" data-line-end="10">✨Magic ✨</li>
</ul>
<h2 class="code-line" data-line-start=10 data-line-end=11 ><a id="Features_10"></a>Features</h2>
<ul>
<li class="has-line-data" data-line-start="12" data-line-end="13">Pipeline execution from Jenkins</li>
<li class="has-line-data" data-line-start="13" data-line-end="14">Compiling native code in Go</li>
<li class="has-line-data" data-line-start="14" data-line-end="15">Code scanning with sonarqube</li>
<li class="has-line-data" data-line-start="15" data-line-end="16">Compiling and pushing a real docker image to the dockerhub</li>
<li class="has-line-data" data-line-start="16" data-line-end="17">Creation of an application that creates ascii-art from real images.</li>
</ul>
<h2 class="code-line" data-line-start=20 data-line-end=21 ><a id="Tech_20"></a>Tech</h2>
<p class="has-line-data" data-line-start="22" data-line-end="23">Dillinger uses a number of open source projects to work properly:</p>
<ul>
<li class="has-line-data" data-line-start="24" data-line-end="25">[Go] - Language in which pipelines and applications are written.</li>
<li class="has-line-data" data-line-start="25" data-line-end="26">[yaml] - Creation cluster kind</li>
<li class="has-line-data" data-line-start="26" data-line-end="28">[groovy] - Pipelines in Jenkins</li>
</ul>
<h2 class="code-line" data-line-start=28 data-line-end=29 ><a id="Installation_28"></a>Installation</h2>
<p class="has-line-data" data-line-start="30" data-line-end="32">It is necessary to understand that to run the lab correctly you must create a kubernetes cluster immediately after connecting it to Jenkins, guaranteeing access to the dagger agent on the docker sock.<br>
To run each of the pipelines locally you need to execute the following commands:</p>
<pre><code class="has-line-data" data-line-start="34" data-line-end="38" class="language-sh">go get dagger.io/dagger@latest
go mod tidy
go run script.go
</code></pre>
<p class="has-line-data" data-line-start="39" data-line-end="40">Kind-configuration</p>
<ul>
<li class="has-line-data" data-line-start="40" data-line-end="42">You will find the basic configuration to create a kind cluster.</li>
</ul>
<p class="has-line-data" data-line-start="42" data-line-end="43">jenkins-agent-dagger</p>
<ul>
<li class="has-line-data" data-line-start="44" data-line-end="46">Inside you will find the base dockerfile to create your dagger agent.</li>
</ul>
<p class="has-line-data" data-line-start="46" data-line-end="47">initial-demo</p>
<ul>
<li class="has-line-data" data-line-start="48" data-line-end="50">You will find a simple script where the basic structure of a dagger pipeline is run</li>
</ul>
<p class="has-line-data" data-line-start="50" data-line-end="51">continous-integration</p>
<ul>
<li class="has-line-data" data-line-start="52" data-line-end="54">You will find a more complex structure where the scanning of the same pipeline is done with sonarqube.</li>
</ul>
<p class="has-line-data" data-line-start="54" data-line-end="55">ascii-art</p>
<ul>
<li class="has-line-data" data-line-start="56" data-line-end="58">This is the most complex of all where you will find a go application that will be compiled, scanned and containerised.</li>
</ul>
<p class="has-line-data" data-line-start="58" data-line-end="60">For production environments…<br>
In the event that you have to handle secrets, there are several possibilities:</p>
<ul>
<li class="has-line-data" data-line-start="60" data-line-end="61">Mozilla sops</li>
<li class="has-line-data" data-line-start="61" data-line-end="62">Azure Key Vault</li>
<li class="has-line-data" data-line-start="62" data-line-end="64">KMS</li>
</ul>
<p class="has-line-data" data-line-start="64" data-line-end="66">(<a href="https://docs.dagger.io/723462/use-secrets/">https://docs.dagger.io/723462/use-secrets/</a>)<br>
Among others.</p>
<pre><code class="has-line-data" data-line-start="67" data-line-end="71" class="language-sh">go get dagger.io/dagger@latest
go mod tidy
go run script.go
</code></pre>