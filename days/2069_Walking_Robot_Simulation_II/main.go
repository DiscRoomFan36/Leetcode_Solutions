package main


type Robot struct {
    width, height int;

    num_steps int;

    num_steps_was_updated bool;

    cached_x, cached_y int;
    cached_dir string;
}


func Constructor(width int, height int) Robot {
    robot := Robot{
        width: width,
        height: height,
        num_steps: 0,

        num_steps_was_updated: true,
    };
    return robot;
}


func (this *Robot) Step(num int)  {
    this.num_steps += num;
    this.num_steps_was_updated = true;
}

func (this *Robot) Get_Pos_And_Direction() (int, int, string) {
    if this.num_steps_was_updated {
        this.num_steps_was_updated = false;
        x, y, dir := this.get_pos_and_direction_for_real();
        this.cached_x = x;
        this.cached_y = y;
        this.cached_dir = dir;
    }

    return this.cached_x, this.cached_y, this.cached_dir;
}

func (this *Robot) get_pos_and_direction_for_real() (int, int, string) {
    width, height := this.width, this.height;
    num_steps_arount_perimeter := width * 2 + (height - 2) * 2;

    steps := this.num_steps % num_steps_arount_perimeter;

    // special case the 0th step.
    if steps == 0 {
        if this.num_steps == 0 {
            return 0, 0, "East";
        } else {
            return 0, 0, "South";
        }
    }

    if steps < width {
        return steps, 0, "East";
    }
    if steps < width + height-1 {
        return width-1, steps+1 - width, "North";
    }
    if steps < width + height-1 + width-1 {
        return width-1 - (steps+1 - (width + height-1)), height-1, "West";
    }
    if steps < width + height-1 + width-1 + height-1 {
        return 0, height-1 - (steps+1 - (width + height-1 + width-1)), "South";
    }

    panic("UNREACHABLE");
}

func (this *Robot) GetPos() []int {
    x, y, _ := this.Get_Pos_And_Direction();
    return []int{x, y};
}


func (this *Robot) GetDir() string {
    _, _, dir := this.Get_Pos_And_Direction();
    return dir;
}


/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
