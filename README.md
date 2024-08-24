<h1>simple api made in gofiber</h1>

<strong>docker image</strong> - docker pull quaintlotus8443/fiber-app:v1 <br>
<strong>live @</strong> - https://fiber-app-b27i.onrender.com/

<h2>Installation instructions</h2>
<ol>
    <li>Running on bare-metal</li>
        <ul>
            <li>Git clone the repo</li>
            <li>Open project dir and install the dependencies</li>
            <li>Put your database url in the .env file</li>
            <li>go build -o appname</li>
        </ul>
    <li>Running inside a Docker image</li>
        <ul>
            <li>docker buildx builld -t fiber-app .</li>
            <li>docker run -p 3000:3000 fiber-app</li>
        <ul>
</ol>