
docker build --build-arg EXERCISM_TOKEN=$ENV:EXERCISM_TOKEN -t exercism_build dockerfiles/exercism_build

$supported_languages = @("python", "go")

foreach ($lang in $supported_languages) {
    $image_name = "$lang" + "_test"
    $location = "$lang" + "_test_environment"
    docker build -t $image_name dockerfiles/$location
}