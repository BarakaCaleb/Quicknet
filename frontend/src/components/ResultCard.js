// ResultCard.js
const ResultCard = ({ title, value, unit }) => {
    return (
      <div className="p-4 bg-white rounded shadow-md text-center">
        <h3 className="text-lg font-semibold">{title}</h3>
        <p className="text-2xl font-bold mt-2">
          {value} <span className="text-sm">{unit}</span>
        </p>
      </div>
    );
  };
  
  export default ResultCard;
  