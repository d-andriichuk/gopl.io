package lenconv

//Convert length from kilometers to miles
func KMToMile(km Kilometre) Mile { return Mile(km / 1.2) }

//Convert length from miles to kilometers
func MileToKM(mile Mile) Kilometre { return Kilometre(mile * 1.2) }
