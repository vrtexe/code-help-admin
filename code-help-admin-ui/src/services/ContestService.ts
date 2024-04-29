import { ContestApi, type ContestEditRequest, type ContestRequest } from "../generated/admin-api";
import { baseConfiguration } from "./api";

const CONTEST_API = new ContestApi(baseConfiguration);

export const getAllContests = () => CONTEST_API.getAllContests();
export const getContestById = (id: number) => CONTEST_API.getContest({ id });
export const createContest = (body: ContestRequest) => CONTEST_API.createContest({ contestRequest: body });
export const updateContest = (id: number, body: ContestEditRequest) =>
  CONTEST_API.updateContest({ id, contestEditRequest: body });
export const deleteContest = (id: number) => CONTEST_API.deleteContest({ id });
