package models

type Member struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
}

func GetEnhypenMembers() []Member {
    members := []Member{
        {"Jungwon", 18},
        {"Heeseung", 20},
        {"Jay", 18},
        {"Jake", 19},
        {"Sunghoon", 18},
        {"Sunoo", 17},
        {"Ni-Ki", 16},
    }
    return members
}

func GetEnhypenMemberByName(name string) Member {
    members := GetEnhypenMembers()
    for _, member := range members {
        if member.Name == name {
            return member
        }
    }
    return Member{}
}
