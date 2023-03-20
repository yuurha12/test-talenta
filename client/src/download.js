import { API } from "./config/api";
import jsPDF from "jspdf";
import "jspdf-autotable";

export default function downloadPdf() {
  API.get("/friends")
    .then((response) => {
      const friends = response.data.data;
      const doc = new jsPDF();
      const canvas = document.getElementById("chart");
      const chartData = canvas.toDataURL();

      doc.autoTable({
        head: [["Name", "Gender", "Age"]],
        body: friends?.map((friend) => [friend.name, friend.gender, friend.age]),
      });

      doc.addPage();
      doc.addImage(chartData, "PNG", 10, 10, 180, 100);

      API.get("/friendstats")
        .then((response) => {
          doc.autoTable({
            margin: { top: 120 },
            head: [["Total Friend", "Male Count", "Female Count", "Under 19 Count", "Above 20 Count"]],
            body: [[
              response.data.data.total_friend_count,
              response.data.data.male_count,
              response.data.data.female_count,
              response.data.data.under_19_count,
              response.data.data.above_20_count,
            ]],
          });

          doc.save("friends.pdf");
        })
        .catch((error) => console.log(error));
    })
    .catch((error) => console.log(error));
}
