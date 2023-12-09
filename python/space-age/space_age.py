class SpaceAge:
    def __init__(self, seconds):
        self.earth_seconds = seconds / one_earth_year_second
    def on_earth(self):
        return round(self.earth_seconds/planet_to_earth_years["earth"],2)
    def on_mercury(self):
        return round(self.earth_seconds/planet_to_earth_years["mercury"],2)
    def on_venus(self):
        return round(self.earth_seconds/planet_to_earth_years["venus"],2)
    def on_mars(self):
        return round(self.earth_seconds/planet_to_earth_years["mars"],2)
    def on_jupiter(self):
        return round(self.earth_seconds/planet_to_earth_years["jupiter"],2)
    def on_saturn(self):
        return round(self.earth_seconds/planet_to_earth_years["saturn"],2)
    def on_uranus(self):
        return round(self.earth_seconds/planet_to_earth_years["uranus"],2)
    def on_neptune(self):
        return round(self.earth_seconds/planet_to_earth_years["neptune"],2)



planet_to_earth_years = {
	"mercury": 0.2408467,
	"venus":   0.61519726,
	"earth":   1,
	"mars":    1.8808158,
	"jupiter": 11.862615,
	"saturn":  29.447498,
	"uranus":  84.016846,
	"neptune": 164.79132,
}

one_earth_year_second = 31557600
