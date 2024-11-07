class Event{
    event={};
    on(eventName,handler){
        if(!this.event[eventName]){
            this.event[eventName]=[];
        }
        this.event[eventName].push(handler);
    }
    emit(eventName,data){
        if(this.event[eventName]){
            for(const handler of this.event[eventName]){
                handler(data);
            }
        }
    }
}

const emit=new Event();
export default emit;