import { API } from "./config/api";
import jsPDF from "jspdf";
import "jspdf-autotable";

export default function downloadPdf() {
  // Get friend data from API
  API
    .get("/friends")
    .then((response) => {
      const friends = response.data.data;
      // Create PDF document using jsPDF and add table
      const doc = new jsPDF();
      doc.autoTable({
        head: [["Name", "Gender", "Age"]],
        body: friends?.map((friend) => [friend.name, friend.gender, friend.age]),
      });

      // Add chart as image to PDF document
      const canvas = document.getElementById("chart");
      const chartData = canvas.toDataURL();
      doc.addPage();
      doc.addImage(chartData, "PNG", 10, 10, 180, 100);

      // Download PDF file
      doc.save("friends.pdf");
    })
    .catch((error) => console.log(error));
}
