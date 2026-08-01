import Room from "../db/models/room"
import BaseRepository from "./base.repositry"


export class RoomRepository extends BaseRepository<Room>{
    constructor(){
        super(Room);
    }
}