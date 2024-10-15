CREATE TABLE "filmByGenre"(
    "id" SERIAL PRIMARY KEY,
    "genreFilmId" INT REFERENCES "genreFilm"("id"), 
    "filmId" INT REFERENCES "film"("id")
)