package main

// a node in a tree.
type Person struct {
    name string;
    is_alive bool;

    parent *Person;
    // from oldest to youngest.
    children []*Person;
}

type ThroneInheritance struct {
    // pointer for pointer safety reasons.
    king *Person;

    name_to_person map[string]*Person;
}


func Constructor(kingName string) ThroneInheritance {
    king := new(Person);
    king.name = kingName;
    king.is_alive = true;

    name_to_person := make(map[string]*Person);
    name_to_person[kingName] = king;

    result := ThroneInheritance{
        king: king,
        name_to_person: name_to_person,
    };
    return result;
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {
    parent := this.name_to_person[parentName];
    child := new(Person);
    child.name = childName;
    child.is_alive = true;

    child.parent = parent;
    parent.children = append(parent.children, child);

    this.name_to_person[childName] = child;
}


func (this *ThroneInheritance) Death(name string)  {
    person := this.name_to_person[name];
    person.is_alive = false;
}


func (this *ThroneInheritance) GetInheritanceOrder() []string {
    result := make([]string, 0);

    in_list := make(map[*Person]bool);

    var successor func(person *Person) *Person;
    successor = func(person *Person) *Person {
        // Successor(x, curOrder):
        //   if x has no children or all of x's children are in curOrder:
        //       if x is the king return null
        //       else return Successor(x's parent, curOrder)
        //   else return x's oldest child who's not in curOrder

        for _, child := range person.children {
            if in_list[child] { continue; }
            return child;
        }

        // if you are the king.
        if person.parent == nil { return nil; }
        return successor(person.parent);
    }

    current_person := this.king;
    for current_person != nil {
        if current_person.is_alive {
            result = append(result, current_person.name);
        }
        in_list[current_person] = true;

        // get the next person
        current_person = successor(current_person);
    }
    return result;
}


/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
