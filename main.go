package main

import "vistar_cacher/downloader"


func main() {
	// 1. Create the cache
	// 2. Download images
	// 3. Store image in a file()

	our_image := []string{
		"https://www.bluecross.org.uk/sites/default/files/assets/images/124044lpr.jpg",
		"https://static.boredpanda.com/blog/wp-content/uploads/2018/04/5acb63d83493f__700-png.jpg",
		"https://www.argospetinsurance.co.uk/assets/uploads/2017/12/cat-pet-animal-domestic-104827.jpeg",
		"https://www.pets4homes.co.uk/images/articles/771/large/cat-lifespan-the-life-expectancy-of-cats-568e40723c336.jpg",
	}

	for _, image := range our_image {
		downloader.ImageDownload(image)
	}


	

}
