package main

import (
	"cmp"
	"slices"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    n := len(positions);

    type Robot struct {
        id int;
        position int;
        health int;
        going_right bool;
    };

    robots := make([]Robot, n);
    for i := range n {
        robots[i] = Robot{
            id: i+1,
            position:    positions[i],
            health:      healths[i],
            going_right: directions[i] == 'R',
        };
    }

    // sort by position.
    slices.SortFunc(robots,
        func(a, b Robot) int {
            return cmp.Compare(a.position, b.position);
        },
    );

    // all the robots that cannot hit anything
    free_robots := make([]Robot, 0);

    // get starting robots going left, no opposition.
    for (len(robots) > 0 && !robots[0].going_right) {
        Append(&free_robots, robots[0]);
        // pop front.
        robots = robots[1:];
    }

    // get ending robots going right, no opposition
    for (len(robots) > 0 && robots[len(robots)-1].going_right) {
        Append(&free_robots, robots[len(robots)-1]);
        robots = robots[:len(robots)-1];
    }


    // for efficently removeing robots from the list.
    dl := Doubly_Linked_List[Robot]{};
    for _, r := range robots { dl.add(r); }

    for dl.count > 0 {
        first_going_left := dl.head;
        for ; first_going_left != nil; first_going_left = first_going_left.next {
            if first_going_left.item.health <= 0 { panic("what"); }
            if !first_going_left.item.going_right {
                break;
            }
        }

        if first_going_left == nil {
            for node := dl.head; node != nil; node = node.next {
                Append(&free_robots, node.item);
            }
            dl.clear();
            break;
        }

        // first_going_left == dl.head, is fine. its handled correctly.
        going_left_robot := first_going_left.item;
        robot_to_the_left := first_going_left.prev;

        // this node is no longer useful
        dl.remove(first_going_left);

        // now go left, and check what this will destroy
        for robot_to_the_left != nil {
            this_robot := &robot_to_the_left.item;

            if this_robot.health < going_left_robot.health {
                robot_to_the_left, _ = dl.remove(robot_to_the_left);
                going_left_robot.health -= 1;
            } else if this_robot.health > going_left_robot.health {
                this_robot.health -= 1;
                going_left_robot.health = 0;
                break;
            } else {
                // health is equal
                dl.remove(robot_to_the_left);
                going_left_robot.health = 0;
                break;
            }
        }

        if going_left_robot.health > 0 { Append(&free_robots, going_left_robot); }
    }


    // sort free robots by id.
    slices.SortFunc(free_robots,
        func(a, b Robot) int {
            return cmp.Compare(a.id, b.id);
        },
    );
    result := make([]int, len(free_robots));
    for i, r := range free_robots {
        result[i] = r.health;
    }
    return result;
}


type Doubly_Linked_List[T any] struct {
    head, tail *DL_Node[T];
    count int;
}

type DL_Node[T any] struct {
    next, prev *DL_Node[T];
    item T;
}

func (dl *Doubly_Linked_List[T]) clear() {
    dl.head  = nil;
    dl.tail  = nil;
    dl.count = 0;
}

func (dl *Doubly_Linked_List[T]) add(new_item T) {
    dl.count += 1;

    new_node := new(DL_Node[T]);
    new_node.item = new_item;

    if dl.head == nil {
        // empty dl
        if dl.tail != nil { panic("what"); }
        dl.head = new_node;
        dl.tail = new_node;

    } else {
        new_node.prev = dl.tail;
        dl.tail.next = new_node;

        dl.tail = new_node;
    }
}

func (dl *Doubly_Linked_List[T]) remove(node *DL_Node[T]) (left_node, right_node *DL_Node[T]) {
    if node == nil { panic("why did you pass this to us."); }
    dl.count -= 1;

    if dl.head == node && dl.tail == node {
        dl.head = nil;
        dl.tail = nil;
        return nil, nil;
    }
    if dl.head == node {
        dl.head.next.prev = nil;
        dl.head = dl.head.next;
        return nil, dl.head;
    }
    if dl.tail == node {
        dl.tail.prev.next = nil;
        dl.tail = dl.tail.prev;
        return dl.tail, nil;
    }

    node.prev.next = node.next;
    node.next.prev = node.prev;
    return node.prev, node.next;
}



func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}
