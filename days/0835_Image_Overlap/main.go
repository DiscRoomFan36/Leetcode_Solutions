package main

func largestOverlap(img1 [][]int, img2 [][]int) int {
    type Pixel struct { x, y int; };

    n := len(img1);

    get_pixels := func(img [][]int) []Pixel {
        img_pixels := make([]Pixel, 0);

        for j := range n {
            for i := range n {
                if img[j][i] == 1 {
                    Append(&img_pixels, Pixel{i, j});
                }
            }
        }

        return img_pixels;
    }

    img_1_pixels := get_pixels(img1);
    img_2_pixels := get_pixels(img2);

    // make img_1 the smaller image
    if len(img_1_pixels) > len(img_2_pixels) {
        img_1_pixels, img_2_pixels = img_2_pixels, img_1_pixels;
        img1, img2 = img2, img1;
    }

    best_score := 0;
    // loop over all posible transforms
    for j := -n+1; j < n; j += 1 {
        for i := -n+1; i < n; i += 1 {

            // count overlap for this transform
            total_overlap := 0;
            // this is why we nade img_1_pixels the smaller one,
            // so this loop would be smaller.
            for _, pixel := range img_1_pixels {
                x, y := pixel.x + i, pixel.y + j;
                if x < 0 || x >= n { continue; }
                if y < 0 || y >= n { continue; }

                // a pretty fast way to check,
                if img2[y][x] == 1 {
                    total_overlap += 1;
                }
            }

            best_score = max(best_score, total_overlap);
            // if its allready the best possible score, dont continue.
            if best_score == len(img_1_pixels) { break; }
        }
    }

    return best_score;
}

func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
