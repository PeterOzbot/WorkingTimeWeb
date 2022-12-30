export default class Month {
    public name: String;
    public id: Number;
    public selected: boolean;

    constructor(name: string, id: Number) {
        this.name = name;
        this.id = id;
        this.selected = false;
    }
}