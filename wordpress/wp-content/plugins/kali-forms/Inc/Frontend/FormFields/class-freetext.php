<?php
namespace KaliForms\Inc\Frontend\FormFields;

if (!defined('ABSPATH')) {
    exit;
}

/**
 * Class FreeText
 *
 * @package Inc\Frontend\FormFields;
 */
class FreeText extends Form_Field
{
    /**
     * Class constructor
     */
    public function __construct()
    {
        $this->id = 'freeText';
    }

    /**
     * Render function
     *
     * @return void
     */
    public function render($item, $form_info)
    {
        $attributes = $this->generate_attribute_string($item);

        $div = '<div class="col-12 col-md-' . absint($item['col']) . '">';
        $div .= wp_kses_post($item['content']);
        $div .= '</div>';

        return $div;

    }
}
