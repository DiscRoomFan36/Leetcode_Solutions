package main

const MOD = 1000000007;

type Instruction struct {
    // is add or mult.
    is_add bool;
    // number to do that by
    num    int;

    size_of_array_when_made int;
}

type Val_And_Index struct {
    val int;
    instruction_index int;
}

type Fancy struct {
    base_sequence []Val_And_Index;

    instructions []Instruction;

    // when index was called, this is the last instruction
    last_instruciton_when_GetIndex_was_called int;
}


func Constructor() Fancy {
    result := Fancy{
        base_sequence: make([]Val_And_Index, 0),

        instructions: make([]Instruction, 0),

        last_instruciton_when_GetIndex_was_called: -1,
    };
    return result;
}


func (this *Fancy) Append(val int)  {
    Append(&this.base_sequence, Val_And_Index{val, len(this.instructions)});
}

func (this *Fancy) add_instruction(new_instruction Instruction) {
    // dont add anything if there is nothing to add to.
    if new_instruction.size_of_array_when_made == 0 { return; }

    // if there were 2 add instructions in a row, maybe we can fold them.
    maybe_fold_instruction := (len(this.instructions) > 0) && (this.last_instruciton_when_GetIndex_was_called != len(this.instructions));

    // fmt.Println(this, maybe_fold_instruction);

    if maybe_fold_instruction {
        last_instruction := this.instructions[len(this.instructions)-1];

        maybe_fold_instruction = maybe_fold_instruction && (last_instruction.is_add                  == new_instruction.is_add);
        maybe_fold_instruction = maybe_fold_instruction && (last_instruction.size_of_array_when_made == new_instruction.size_of_array_when_made);
    }

    if maybe_fold_instruction {
        if new_instruction.is_add {
            this.instructions[len(this.instructions)-1].num += new_instruction.num;
        } else {
            this.instructions[len(this.instructions)-1].num *= new_instruction.num;
        }
        this.instructions[len(this.instructions)-1].num %= MOD;

    } else {

        Append(&this.instructions, new_instruction);
    }
}

func (this *Fancy) AddAll(inc int)  {
    new_instruction := Instruction{
        is_add: true,
        num: inc,
        size_of_array_when_made: len(this.base_sequence),
    };
    this.add_instruction(new_instruction);
}


func (this *Fancy) MultAll(m int)  {
    new_instruction := Instruction{
        is_add: false,
        num: m,
        size_of_array_when_made: len(this.base_sequence),
    };
    this.add_instruction(new_instruction);
}


func (this *Fancy) GetIndex(idx int) int {
    if idx < 0 || len(this.base_sequence) <= idx { return -1; }

    // fmt.Println(this);

    this.last_instruciton_when_GetIndex_was_called = len(this.instructions);

    val_and_index := &this.base_sequence[idx];

    for ; val_and_index.instruction_index < len(this.instructions); val_and_index.instruction_index++ {
        instruction := this.instructions[val_and_index.instruction_index];

        if instruction.is_add {
            val_and_index.val = (val_and_index.val + instruction.num) % MOD;
        } else {
            val_and_index.val = (val_and_index.val * instruction.num) % MOD;
        }
    }

    return val_and_index.val;
}


func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}


/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */