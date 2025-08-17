import NavBar from "../../components/NavBar";
import Card from "../../components/Card";
import ClickCard from "../../components/ClickCard";
import { useNavigate } from "react-router-dom";
import Heading from "../../components/Heading";
import { MdGroups } from "react-icons/md";
import { FaBook,FaDoorOpen,FaPlus,FaEdit } from "react-icons/fa";
import { FaTableCells } from "react-icons/fa6";

const FacultyDashboard = () => {
  const navigate = useNavigate();

  return (
    <div className="bg-gray-100 h-full">
      <NavBar />
      <Heading text="Faculty Dashboard" />
      <div className="flex justify-baseline gap-3 items-center p-5">
        <Card heading="Heading" title="Faculty Members" icon={MdGroups } iconColor={"bg-green-600"} />
        <Card heading="Heading" title="Subjects" icon={FaBook } iconColor={"bg-orange-600"} />
        <Card heading="Heading" title="Rooms" icon={FaDoorOpen } iconColor={"bg-red-600"} />
      </div>
      <div className="flex flex-col p-5 gap-5 justify-baseline items-center m-2 md:flex-row md:flex-wrap">
        
        <ClickCard
          discription="View Timetables"
          title="View"
          icon={FaTableCells }
          iconColor="bg-green-600"
          onClick={() => navigate("/view-timetable")}
        />

      </div>
    </div>
  );
};

export default FacultyDashboard;
