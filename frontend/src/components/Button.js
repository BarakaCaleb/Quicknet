// Button.js
const Button = ({ onClick, label }) => {
    return (
      <button
        onClick={onClick}
        className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-green-700 transition"
      >
        {label}
      </button>
    );
  };
  
  export default Button;
  